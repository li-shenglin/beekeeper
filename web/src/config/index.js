export default {
  title: 'beekeeper',
  cookieExpires: 30,
  useI18n: true,
  baseUrl: {
    dev: '/apis',
    pro: '/apis'
  },
  homeName: 'home',
  plugin: {
    'error-store': {
      showInHeader: true,
      developmentOff: true
    }
  }
}
