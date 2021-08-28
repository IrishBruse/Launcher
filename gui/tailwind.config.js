module.exports = {
    purge: ['./public/index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    darkMode: false,
    theme: {
        extend: {},
        colors: {
            'primary': {
                dark: "var(--backgroundDark)",
                DEFAULT: "var(--backgroundMedium)",
                light: "var(--backgroundLight)",
            },
            'secondary': 'var(--text)',
            'secondary-hover': 'var(--subText)',
            'tertiary': 'var(--link)',
        }
    },
    fontFamily: {
        sans: ['Open Sans', 'sans-serif'],
    },
    variants: {
        extend: {},
    },
    plugins: [],
}
