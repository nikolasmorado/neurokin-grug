/** @type {import("tailwindcss").Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  theme: {
    extend: {
      fontFamily: {
        "inter": ["Inter", "sans-serif"],
        "display": ["degular-display", "sans-serif"],
        "text": ["degular-text", "sans-serif"],
        "anek": ["Anek Bangla", "sans-serif"],
      },
    },
    colors: {
      white: "#ffffff",
      black: "#000000",
      gray: "#D9D9D9",
      green: {
        "300": "#BAFFEC",
        "500": "#8ED9C5",
        "700": "#79F6D3",
      },
      pink: {
        "300": "#FFEAF6",
        "400": "#F6A7D6",
        "500": "#F5A6D5",
        "700": "#E26FA3",
      },
      red: {
        "500": "#EF4444",
      }
    }
  },
}
