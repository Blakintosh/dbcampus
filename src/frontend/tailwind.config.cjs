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
			'blinking': 'blink 3s step-end infinite',
            'fly-in': 'flyIn 0.25s ease-in'
		},
		keyframes: {
			blink: {
				'0%, 100%': { "background-color": "rgb(245 158 11)" },
				'50%': { "background-color": "black" }
			},
            flyIn: {
                '0%': {
                    "bottom": "-120px",
                    "opacity": "0"
                },
                "100%": {
                    "bottom": "0",
                    "opacity": "1"
                }
            }
		}
	},
  },
  plugins: [],
}
