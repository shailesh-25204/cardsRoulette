/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],  theme: {
    extend: {
      backgroundImage: {
        'main-texture': "url('/src/assets/denim.png')",
      }
    },
  },
  plugins: [],
}

