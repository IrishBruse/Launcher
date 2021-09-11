module.exports = {
    purge: ['./public/index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    theme: {
        extend: {
        },
        colors: {
            'transparent': "transparent",
            'primary': {
                dark: "hsl(203, 12%, 11%)",
                DEFAULT: "hsl(206, 9%, 13%)",
                light: "hsl(206, 9%, 15%)",
                hover: "hsl(209, 9%, 17%)",
            },
            'secondary': {
                dark: 'hsl(214, 5%, 73%)',
                DEFAULT: 'hsl(204, 12%, 92%)',
            },
            'tertiary': {
                dark: 'hsl(85, 63%, 30%)',
                DEFAULT: 'hsl(85, 59%, 48%)',
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
