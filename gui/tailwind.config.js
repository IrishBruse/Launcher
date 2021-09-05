module.exports = {
    purge: ['./public/index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    darkMode: false,
    theme: {
        extend: {
        },
        colors: {
            'primary': {
                dark: "var(--primaryDark)",
                DEFAULT: "var(--primaryMedium)",
                light: "var(--primaryLight)",
                hover: "var(--primaryHover)",
            },
            'secondary': {
                dark: 'var(--secondaryDark)',
                DEFAULT: 'var(--secondaryMedium)',
            },
            'tertiary': {
                dark: 'var(--tertiaryDark)',
                DEFAULT: 'var(--tertiary)',
            },
            'white': 'white',
            'black': 'black',
        }
    },
    fontFamily: {
        sans: ['Open Sans', 'sans-serif'],
    },
    variants: {
        extend: {
            borderRadius: ['hover', 'group-hover', 'focus'],
        }
    },
    plugins: [],
}
