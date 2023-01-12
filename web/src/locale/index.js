import { createI18n } from 'vue-i18n'
import { localRead } from '@/libs/util'
import zh from 'view-ui-plus/dist/locale/zh-CN'
import en from 'view-ui-plus/dist/locale/en-US'
import zhCnLocale from '@/locale/lang/zh-CN'
import enUsLocale from '@/locale/lang/en-US'

const navLang = navigator.language
const localLang = (navLang === 'zh-CN' || navLang === 'en-US') ? navLang : false
let lang = localRead('local') || localLang || 'zh-CN'

const i18n = createI18n({
  allowComposition: true,
  globalInjection: true,
  legacy: false,
  locale: lang,
  messages: {
    'zh-CN': Object.assign(zh, zhCnLocale),
    'en-US': Object.assign(en, enUsLocale)
  }
});
export default i18n

