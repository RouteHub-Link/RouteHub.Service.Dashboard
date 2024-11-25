/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './web/templates/**/*.{js,jsx,ts,tsx,templ,go,txt}',
    './web/templates/**/**/*.{js,jsx,ts,tsx,templ,go,txt}',
    './web/templates/**/**/**/*.{js,jsx,ts,tsx,templ,go,txt}',
    './web/templates/layouts/*.{js,jsx,ts,tsx,templ,go,txt}',
    './web/templates/pages/*.{js,jsx,ts,tsx,templ,go,txt}',
    './web/templates/pages/**/*.{js,jsx,ts,tsx,templ,go,txt}',
    './web/templates/misc/*.{js,jsx,ts,tsx,templ,go,txt}',
    './web/templates/components/**/*.{js,jsx,ts,tsx,templ,go,txt}',
    './node_modules/preline/dist/*.js',],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@popperjs/core'),
    require('preline/plugin'),
  ],
}

