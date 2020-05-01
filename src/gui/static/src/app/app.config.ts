export const AppConfig = {
  otcEnabled: false,
  maxHardwareWalletAddresses: 1,
  urlForHwWalletVersionChecking: 'https://version.amolecoin.com/amolewallet/version.txt',
  hwWalletDownloadUrlAndPrefix: 'https://downloads.amolecoin.com/amolewallet/amolewallet-firmware-v',
  hwWalletDaemonDownloadUrl: 'https://www.amolecoin.com/downloads/',

  urlForVersionChecking: 'https://version.amolecoin.com/amolecoin/version.txt',
  walletDownloadUrl: 'https://www.amolecoin.com/downloads/',

  priceApiId: 'amole-amolecoin',

  /**
   * This wallet uses the Amolecoin URI Specification (based on BIP-21) when creating QR codes and
   * requesting coins. This variable defines the prefix that will be used for creating QR codes
   * and URLs. IT MUST BE UNIQUE FOR EACH COIN.
   */
  uriSpecificatioPrefix: 'amolecoin',

  languages: [{
      code: 'en',
      name: 'English',
      iconName: 'en.png',
    },
    {
      code: 'zh',
      name: '中文',
      iconName: 'zh.png',
    },
    {
      code: 'es',
      name: 'Español',
      iconName: 'es.png',
    },
  ],
  defaultLanguage: 'en',

  mediumModalWidth: '566px',
};
