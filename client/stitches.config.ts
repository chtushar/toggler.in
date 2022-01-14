import { createStitches } from '@stitches/react';
import { gray } from '@radix-ui/colors';

export const {
  styled,
  css,
  globalCss,
  keyframes,
  getCssText,
  theme,
  createTheme,
  config,
} = createStitches({
  theme: {
    fonts: {
      sans: "'Noto Sans Display', sans-serif",
    },
    fontWeights: {
      bold: 700,
      semiBold: 600,
      medium: 500,
      regular: 400,
    },
    colors: {
      ...gray,
    },
    space: {
      1: '4px',
      2: '8px',
      3: '12px',
      4: '16px',
      5: '20px',
      6: '24px',
      7: '28px',
      8: '32px',
      9: '36px',
      10: '40px',
    },
    fontSizes: {
      9: '9px',
      10: '10px',
      11: '11px',
      12: '12px',
      13: '13px',
      14: '14px',
      18: '18px',
      24: '24px',
      36: '36px',
      48: '48px',
      64: '64px',
      72: '72px',
      96: '96px',
      144: '144px',
      288: '288px',
    },
  },
  media: {
    bp1: '(max-width: 1024px)',
    bp2: '(max-width: 768px)',
    bp3: '(max-width: 640px)',
  },
  utils: {
    paddingX: (value: string) => ({
      paddingLeft: value,
      paddingRight: value,
    }),
    paddingY: (value: string) => ({
      paddingTop: value,
      paddingBottom: value,
    }),
  },
});
