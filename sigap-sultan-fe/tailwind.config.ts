import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/**/*.{js,jsx,ts,tsx}",
    "./node_modules/react-tailwindcss-datepicker/dist/index.esm.js",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
        "top-navbar": "url(/background-image/bg_header_navigation.svg)",
        "bottom-navbar": "url(/background-image/bg_bottom_navigation.svg)"
      },
      fontFamily: {
        sans: ['"Plus Jakarta Sans"', 'sans-serif'],
      },
      colors: {
        "custom-blue": '#005395',
      },
      textColor: {
        "custom-blue": '#005395',
        "custom-yellow": '#F4CC22',
        "custom-orange": '#DC6803',
        "custom-green": '#23A65F',
        "secondary-red": '#B11016'
      }
    },
  },
  plugins: [],
};
export default config;
