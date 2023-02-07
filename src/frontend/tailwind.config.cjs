/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
	fontFamily: {
		'sans': ['Open Sans', 'sans-serif'],
		'serif': ['Arvo', 'serif'],
	},
    extend: {
		animation: {
			'blinking': 'blink 3s step-end infinite'
		},
		keyframes: {
			blink: {
				'0%, 100%': { "background-color": "rgb(245 158 11)" },
				'50%': { "background-color": "black" }
			}
		}
	},
  },
  plugins: [],
}
