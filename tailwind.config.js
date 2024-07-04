const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      spacing: {
        '128': '32rem',
      },
      fontFamily: {
        sans: [
          'BMWTypeNext-Light',
          'Tacoma',
          'Arial',
          'Helvetica',
          ...defaultTheme.fontFamily.sans,
        ],
      },
      colors: {
        primary: {
          500: '#1C69D4',
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}