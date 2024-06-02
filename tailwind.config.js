/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  theme: {
      colors: {
        white: "#ffffff",
        black: "#000000",
        green: {
          "300": "#BAFFEC",
          "500": "#8ED9C5",
          "700": "#79F6D3",
        },
        pink: {
          "300": "#FFEAF6",
          "500": "#F5A6D5",
          "700": "#E26FA3",
        }
      }
  },
}
