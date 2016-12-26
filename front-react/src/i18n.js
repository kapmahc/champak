import i18next from 'i18next';
import XHR from 'i18next-xhr-backend';
import LanguageDetector from 'i18next-browser-languagedetector';

import {LOCALE} from './constants'
import {api} from './ajax'

const main = (func) => {
  i18next
    .use(XHR)
    .use(LanguageDetector)
    .init({
      backend: {
        crossDomain: true,
        loadPath: api('/locales/{{lng}}'),
        // withCredentials: true,
      },
      detection: {
        order: ['querystring', 'cookie', 'localStorage', 'navigator', 'htmlTag'],
        caches: ['localStorage', 'cookie'],
        lookupQuerystring: LOCALE,
        lookupCookie: LOCALE,
        lookupLocalStorage: LOCALE,
        cookieMinutes: 60*24*30,
      }
    },
    (err, t) => {
    // initialized and ready to go!
    // const hw = i18next.t('key'); // hw = 'hello world'
    func();
    }
  );
}

export default main
