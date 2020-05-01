import { SubscriptionLike } from 'rxjs';
import { first } from 'rxjs/operators';
import { Component, EventEmitter, Input, OnDestroy, OnInit, ViewChild, ChangeDetectorRef, Output as AgularOutput } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { BigNumber } from 'bignumber.js';
import { TranslateService } from '@ngx-translate/core';

import { PasswordDialogComponent } from '../../../layout/password-dialog/password-dialog.component';
import { ButtonComponent } from '../../../layout/button/button.component';
import { NavBarSwitchService } from '../../../../services/nav-bar-switch.service';
import { SelectAddressComponent } from '../../../layout/select-address/select-address.component';
import { BlockchainService } from '../../../../services/blockchain.service';
import { HwWalletService } from '../../../../services/hw-wallet.service';
import { ChangeNoteComponent } from '../send-preview/transaction-info/change-note/change-note.component';
import { MsgBarService } from '../../../../services/msg-bar.service';
import { MultipleDestinationsDialogComponent } from '../../../layout/multiple-destinations-dialog/multiple-destinations-dialog.component';
import { FormSourceSelectionComponent, AvailableBalanceData, SelectedSources, SourceSelectionModes } from '../form-parts/form-source-selection/form-source-selection.component';
import { FormDestinationComponent, Destination } from '../form-parts/form-destination/form-destination.component';
import { CopyRawTxComponent, CopyRawTxData } from '../offline-dialogs/implementations/copy-raw-tx.component';
import { DoubleButtonActive } from '../../../../components/layout/double-button/double-button.component';
import { ConfirmationParams, DefaultConfirmationButtons, ConfirmationComponent } from '../../../../components/layout/confirmation/confirmation.component';
import { AppService } from '../../../../services/app.service';
import { SpendingService, HoursDistributionOptions, HoursDistributionTypes } from '../../../../services/wallet-operations/spending.service';
import { GeneratedTransaction, Output } from '../../../../services/wallet-operations/transaction-objects';
import { WalletWithBalance, AddressWithBalance } from '../../../../services/wallet-operations/wallet-objects';

/**
 * Data returned when SendCoinsFormComponent asks to show the preview of a transaction. Useful
 * for showing a preview and for restoring the state of the form.
 */
export interface SendCoinsData {
  /**
   * Data entered on the form.
   */
  form: FormData;
  /**
   * How many coins is the user trying to send.
   */
  amount: BigNumber;
  /**
   * List of all the destination addresses.
   */
  to: string[];
  /**
   * Unsigned transaction which was created and the user wants to preview.
   */
  transaction: GeneratedTransaction;
  /**
   * If true, the transaction is a manually created unsigned transaction which is not mean to be
   * sent to the network. The raw transaction text must be shown to the user, so it can be
   * signed and sent later.
   */
  showForManualUnsigned: boolean;
}

/**
 * Contents of a send coins form.
 */
export interface FormData {
  wallet: WalletWithBalance;
  addresses: AddressWithBalance[];
  /**
   * Addresses the user entered manually. Used when manually creating an unsigned transaction,
   * so there are no fields for selecting a wallet or addresses.
   */
  manualAddresses: string[];
  changeAddress: string;
  destinations: Destination[];
  hoursSelection: HoursDistributionOptions;
  /**
   * If true, the hours must be distributed automatically. If false, the user must manually
   * enter the hours for each destination.
   */
  autoOptions: boolean;
  /**
   * All unspent outputs obtained from the node, not the selected ones.
   */
  allUnspentOutputs: Output[];
  outputs: Output[];
  /**
   * Button selected for choosing which currency to use for the amounts.
   */
  currency: DoubleButtonActive;
  note: string;
}

/**
 * Form for sending coins.
 */
@Component({
  selector: 'app-send-coins-form',
  templateUrl: './send-coins-form.component.html',
  styleUrls: ['./send-coins-form.component.scss'],
})
export class SendCoinsFormComponent implements OnInit, OnDestroy {
  // Default factor used for automatically distributing the coins.
  private readonly defaultAutoShareValue = '0.5';

  // Subform for selecting the sources.
  @ViewChild('formSourceSelection', { static: false }) formSourceSelection: FormSourceSelectionComponent;
  // Subform for entering the destinations.
  @ViewChild('formMultipleDestinations', { static: false }) formMultipleDestinations: FormDestinationComponent;
  @ViewChild('previewButton', { static: false }) previewButton: ButtonComponent;
  @ViewChild('sendButton', { static: false }) sendButton: ButtonComponent;
  // Data the form must have just after being created.
  @Input() formData: SendCoinsData;
  // If true, the simple form will be used.
  @Input() showSimpleForm: boolean;
  // Event emited when the transaction has been created and the user wants to see a preview.
  @AgularOutput() onFormSubmitted = new EventEmitter<SendCoinsData>();

  sourceSelectionModes = SourceSelectionModes;
  doubleButtonActive = DoubleButtonActive;

  // Max chars the note field can have.
  maxNoteChars = ChangeNoteComponent.MAX_NOTE_CHARS;
  form: FormGroup;
  // How many coins the user can send with the selected sources.
  availableBalance = new AvailableBalanceData();
  // If true, the hours are distributed automatically. If false, the user can manually
  // enter how many hours to send to each destination.
  autoHours = true;
  // If true, the options for selecting the auto hours distribution factor are shown.
  autoOptions = false;
  // Factor used for automatically distributing the coins.
  autoShareValue = this.defaultAutoShareValue;
  // If true, the form is shown deactivated.
  busy = false;
  // If true, the form is used for manually creating unsigned transactions.
  showForManualUnsigned = false;
  // Sources the user has selected.
  private selectedSources: SelectedSources;

  private syncCheckSubscription: SubscriptionLike;
  private processingSubscription: SubscriptionLike;

  constructor(
    public appService: AppService,
    private blockchainService: BlockchainService,
    private dialog: MatDialog,
    private msgBarService: MsgBarService,
    private navBarSwitchService: NavBarSwitchService,
    private hwWalletService: HwWalletService,
    private translate: TranslateService,
    private changeDetector: ChangeDetectorRef,
    private spendingService: SpendingService,
  ) { }

  ngOnInit() {
    this.form = new FormGroup({}, this.validateForm.bind(this));
    this.form.addControl('changeAddress', new FormControl(''));
    this.form.addControl('note', new FormControl(''));

    if (this.formData) {
      setTimeout(() => this.fillForm());
    }
  }

  ngOnDestroy() {
    if (this.processingSubscription && !this.processingSubscription.closed) {
      this.processingSubscription.unsubscribe();
    }
    this.closeSyncCheckSubscription();
    this.msgBarService.hide();
  }

  // Called when there are changes in the source selection form.
  sourceSelectionChanged() {
    this.selectedSources = this.formSourceSelection.selectedSources;
    this.availableBalance = this.formSourceSelection.availableBalance;
    this.formMultipleDestinations.updateValuesAndValidity();
    this.form.updateValueAndValidity();
  }

  // Called when there are changes in the destinations form.
  destinationsChanged() {
    setTimeout(() => {
      this.form.updateValueAndValidity();
    });
  }

  // Starts the process for creating a transaction for previewing it.
  preview() {
    this.checkBeforeCreatingTx(true);
    this.changeDetector.detectChanges();
  }

  // Starts the process for creating a transaction for sending it without preview.
  send() {
    this.checkBeforeCreatingTx(false);
  }

  // Chages the mode of the advanced form. The form can be in normal mode and a special
  // mode for manually creating unsigned transactions.
  changeFormType(value: DoubleButtonActive) {
    if ((value === DoubleButtonActive.LeftButton && !this.showForManualUnsigned) || (value === DoubleButtonActive.RightButton && this.showForManualUnsigned)) {
      return;
    }

    if (value === DoubleButtonActive.RightButton) {
      // Ask for confirmation before activating the manual unsigned tx mode.
      const confirmationParams: ConfirmationParams = {
        text: 'send.unsigned-confirmation',
        defaultButtons: DefaultConfirmationButtons.YesNo,
      };

      ConfirmationComponent.openDialog(this.dialog, confirmationParams).afterClosed().subscribe(confirmationResult => {
        if (confirmationResult) {
          this.showForManualUnsigned = true;
        }
      });
    } else {
      this.showForManualUnsigned = false;
    }
  }

  // Sets the factor that will be used for distributing the hours.
  setShareValue(event) {
    this.autoShareValue = parseFloat(event.value).toFixed(2);
  }

  // Opens a modal window for selecting the change address.
  selectChangeAddress() {
    SelectAddressComponent.openDialog(this.dialog).afterClosed().subscribe(response => {
      if (response) {
        this.form.get('changeAddress').setValue(response);
      }
    });
  }

  // Opens the bulk sending modal window with the data the user already added to the form.
  openMultipleDestinationsPopup() {
    let currentString = '';

    // Create a string with the data the user has already entered, using the format of the
    // bulk sending modal window.
    const currentDestinations = this.formMultipleDestinations.getDestinations(false);
    currentDestinations.map(destControl => {
      // Ignore the destinations with no data.
      if (destControl.address.trim().length > 0 ||
        destControl.originalAmount.trim().length > 0 ||
        (!this.autoHours && destControl.hours.trim().length > 0)) {
          // Add the data without potentially problematic characters.
          currentString += destControl.address.replace(',', '');
          currentString += ', ' + destControl.originalAmount.replace(',', '');
          if (!this.autoHours) {
            currentString += ', ' + destControl.hours.replace(',', '');
          }
          currentString += '\r\n';
      }
    });

    MultipleDestinationsDialogComponent.openDialog(this.dialog, currentString).afterClosed().subscribe((response: Destination[]) => {
      if (response) {
        if (response.length > 0) {
          // If the first destination does not have hours, no destination has hours.
          this.autoHours = response[0].hours === undefined;
          setTimeout(() => this.formMultipleDestinations.setDestinations(response));
        } else {
          this.formMultipleDestinations.resetForm();
        }
      }
    });
  }

  // Shows or hides the hours distribution options.
  toggleOptions(event) {
    event.stopPropagation();
    event.preventDefault();

    // Resets the hours distribution options.
    this.autoShareValue = this.defaultAutoShareValue;

    this.autoOptions = !this.autoOptions;
  }

  // Activates/deactivates the option for automatic hours distribution.
  setAutoHours(event) {
    this.autoHours = event.checked;
    this.formMultipleDestinations.updateValuesAndValidity();

    if (!this.autoHours) {
      this.autoOptions = false;
    }
  }

  // Fills the form with the provided values.
  private fillForm() {
    this.showForManualUnsigned = this.formData.showForManualUnsigned,

    this.formSourceSelection.fill(this.formData);
    this.formMultipleDestinations.fill(this.formData);

    ['changeAddress', 'note'].forEach(name => {
      this.form.get(name).setValue(this.formData.form[name]);
    });

    if (this.formData.form.hoursSelection.type === HoursDistributionTypes.Auto) {
      this.autoShareValue = this.formData.form.hoursSelection.share_factor;
      this.autoHours = true;
    } else {
      this.autoHours = false;
    }

    this.autoOptions = this.formData.form.autoOptions;
  }

  // Validates the form.
  private validateForm() {
    if (!this.form) {
      return { Required: true };
    }

    // Check the validity of the subforms.
    if (!this.formSourceSelection || !this.formSourceSelection.valid || !this.formMultipleDestinations || !this.formMultipleDestinations.valid) {
      return { Invalid: true };
    }

    return null;
  }

  // Checks if the blockchain is synchronized. It continues normally creating the tx if the
  // blockchain is synchronized and asks for confirmation if it is not. It does nothing if
  // the form is not valid or busy.
  private checkBeforeCreatingTx(creatingPreviewTx: boolean) {
    if (!this.form.valid || this.previewButton.isLoading() || this.sendButton.isLoading()) {
      return;
    }

    this.closeSyncCheckSubscription();
    this.syncCheckSubscription = this.blockchainService.progress.pipe(first()).subscribe(response => {
      if (response.synchronized) {
        this.prepareTransaction(creatingPreviewTx);
      } else {
        const confirmationParams: ConfirmationParams = {
          text: 'send.synchronizing-warning',
          defaultButtons: DefaultConfirmationButtons.YesNo,
        };

        ConfirmationComponent.openDialog(this.dialog, confirmationParams).afterClosed().subscribe(confirmationResult => {
          if (confirmationResult) {
            this.prepareTransaction(creatingPreviewTx);
          }
        });
      }
    });
  }

  // Makes the preparation steps, like asking for the password, and then calls the function
  // for creating the transaction.
  private prepareTransaction(creatingPreviewTx: boolean) {
    this.msgBarService.hide();
    this.previewButton.resetState();
    this.sendButton.resetState();

    // Request the password only if the wallet is encrypted and the transaction is going
    // to be sent without preview.
    if (!this.showForManualUnsigned && this.selectedSources.wallet.encrypted && !this.selectedSources.wallet.isHardware && !creatingPreviewTx) {
      PasswordDialogComponent.openDialog(this.dialog, { wallet: this.selectedSources.wallet }).componentInstance.passwordSubmit
        .subscribe(passwordDialog => {
          this.createTransaction(creatingPreviewTx, passwordDialog);
        });
    } else {
      if (creatingPreviewTx || this.showForManualUnsigned || !this.selectedSources.wallet.isHardware) {
        this.createTransaction(creatingPreviewTx);
      } else {
        // If using a hw wallet, check the device first.
        this.showBusy(creatingPreviewTx);
        this.processingSubscription = this.hwWalletService.checkIfCorrectHwConnected(this.selectedSources.wallet.addresses[0].address).subscribe(
          () => this.createTransaction(creatingPreviewTx),
          err => this.showError(err),
        );
      }
    }
  }

  // Creates a transaction with the data entered on the form.
  private createTransaction(creatingPreviewTx: boolean, passwordDialog?: any) {
    this.showBusy(creatingPreviewTx);

    // Process the source addresses.
    let selectedAddresses: string[];
    if (!this.showForManualUnsigned) {
      selectedAddresses = this.selectedSources.addresses && this.selectedSources.addresses.length > 0 ?
        this.selectedSources.addresses.map(addr => addr.address) : null;
    } else {
      selectedAddresses = this.selectedSources.manualAddresses;
    }

    // Process the source outputs.
    const selectedOutputs = this.selectedSources.unspentOutputs && this.selectedSources.unspentOutputs.length > 0 ?
      this.selectedSources.unspentOutputs.map(addr => addr.hash) : null;

    const destinations = this.formMultipleDestinations.getDestinations(true);

    // Create the transaction.
    this.processingSubscription = this.spendingService.createTransaction(
      this.selectedSources.wallet,
      selectedAddresses ? selectedAddresses : this.selectedSources.wallet.addresses.map(address => address.address),
      selectedOutputs,
      destinations,
      this.hoursSelection,
      this.form.get('changeAddress').value ? this.form.get('changeAddress').value : null,
      passwordDialog ? passwordDialog.password : null,
      creatingPreviewTx || this.showForManualUnsigned,
    ).subscribe(transaction => {
      // Close the password dialog, if it exists.
      if (passwordDialog) {
        passwordDialog.close();
      }

      const note = this.form.value.note.trim();
      transaction.note = note;

      if (!creatingPreviewTx) {
        if (!this.showForManualUnsigned) {
          // Send the transaction to the network.
          this.processingSubscription = this.spendingService.injectTransaction(transaction.encoded, note)
            .subscribe(noteSaved => {
              let showDone = true;
              // Show a warning if the transaction was sent but the note was not saved.
              if (note && !noteSaved) {
                this.msgBarService.showWarning(this.translate.instant('send.saving-note-error'));
                showDone = false;
              }

              this.showSuccess(showDone);
            }, error => this.showError(error));
        } else {
          const data: CopyRawTxData = {
            rawTx: transaction.encoded,
            isUnsigned: true,
          };

          // Show the raw transaction.
          CopyRawTxComponent.openDialog(this.dialog, data).afterClosed().subscribe(() => {
            this.resetState();

            const confirmationParams: ConfirmationParams = {
              text: 'offline-transactions.copy-tx.reset-confirmation',
              defaultButtons: DefaultConfirmationButtons.YesNo,
            };

            // Ask the user if the form should be cleaned, to be able to create a new transaction.
            ConfirmationComponent.openDialog(this.dialog, confirmationParams).afterClosed().subscribe(confirmationResult => {
              if (confirmationResult) {
                this.resetForm();
                this.msgBarService.showDone('offline-transactions.copy-tx.reset-done', 4000);
              }
            });
          });
        }
      } else {
        // Create an object with the form data and emit an event for opening the preview.
        let amount = new BigNumber('0');
        destinations.map(destination => amount = amount.plus(destination.coins));
        this.onFormSubmitted.emit({
          form: {
            wallet: this.selectedSources.wallet,
            addresses: this.selectedSources.addresses,
            manualAddresses: this.selectedSources.manualAddresses,
            changeAddress: this.form.get('changeAddress').value,
            destinations: destinations,
            hoursSelection: this.hoursSelection,
            autoOptions: this.autoOptions,
            allUnspentOutputs: this.formSourceSelection.unspentOutputsList,
            outputs: this.selectedSources.unspentOutputs,
            currency: this.formMultipleDestinations.currentlySelectedCurrency,
            note: note,
          },
          amount: amount,
          to: destinations.map(d => d.address),
          transaction,
          showForManualUnsigned: this.showForManualUnsigned,
        });
        this.busy = false;
        this.navBarSwitchService.enableSwitch();
      }
    }, error => {
      if (passwordDialog) {
        passwordDialog.error(error);
      }

      this.showError(error);
    });
  }

  private resetForm() {
    this.formSourceSelection.resetForm();
    this.formMultipleDestinations.resetForm();
    this.form.get('changeAddress').setValue('');
    this.form.get('note').setValue('');
    this.autoHours = true;
    this.autoOptions = false;
    this.autoShareValue = this.defaultAutoShareValue;
  }

  // Returns the hours distribution options selected on the form, but with the format needed
  // for creating the transaction using the node.
  private get hoursSelection(): HoursDistributionOptions {
    let hoursSelection: HoursDistributionOptions = {
      type: HoursDistributionTypes.Manual,
    };

    if (this.autoHours) {
      hoursSelection = <HoursDistributionOptions> {
        type: HoursDistributionTypes.Auto,
        mode: 'share',
        share_factor: this.autoShareValue,
      };
    }

    return hoursSelection;
  }

  private closeSyncCheckSubscription() {
    if (this.syncCheckSubscription) {
      this.syncCheckSubscription.unsubscribe();
    }
  }

  // Makes the UI to be shown busy and disables the navbar switch.
  private showBusy(creatingPreviewTx: boolean) {
    if (creatingPreviewTx) {
      this.previewButton.setLoading();
      this.sendButton.setDisabled();
    } else {
      this.sendButton.setLoading();
      this.previewButton.setDisabled();
    }
    this.busy = true;
    this.navBarSwitchService.disableSwitch();
  }

  // Cleans the form, stops showing the UI busy, reactivates the navbar switch and, if showDone
  // is true, shows a msg confirming that the transaction has been sent.
  private showSuccess(showDone: boolean) {
    this.busy = false;
    this.navBarSwitchService.enableSwitch();
    this.resetForm();

    if (showDone) {
      this.msgBarService.showDone('send.sent');
      this.sendButton.resetState();
    } else {
      this.sendButton.setSuccess();
      setTimeout(() => {
        this.sendButton.resetState();
      }, 3000);
    }
  }

  // Stops showing the UI busy, reactivates the navbar switch and shows the error msg.
  private showError(error) {
    this.busy = false;
    this.msgBarService.showError(error);
    this.navBarSwitchService.enableSwitch();
    this.previewButton.resetState().setEnabled();
    this.sendButton.resetState().setEnabled();
  }

  // Stops showing the UI busy and reactivates the navbar switch.
  private resetState() {
    this.busy = false;
    this.navBarSwitchService.enableSwitch();
    this.previewButton.resetState().setEnabled();
    this.sendButton.resetState().setEnabled();
  }
}
