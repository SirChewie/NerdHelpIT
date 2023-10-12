/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../templates/**/*.{gohtml,html}","../controllers/**/*.go"],
  theme: {
    extend: {
      colors: {
        "penn-blue": '#000033',
        "neon-blue": '#6366F1',
        "madder": '#A31621',
      }
    },
  },
  plugins: [],
}

