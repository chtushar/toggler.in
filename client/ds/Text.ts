import { styled } from '../stitches.config';

export const Text = styled('span', {
  lineHeight: '1',
  margin: '0',
  fontWeight: 400,
  fontVariantNumeric: 'tabular-nums',
  display: 'block',

  variants: {
    size: {
      9: {
        fontSize: '$9',
      },
      10: {
        fontSize: '$10',
      },
      11: {
        fontSize: '$11',
      },
      12: {
        fontSize: '$12',
      },
      13: {
        fontSize: '$13',
      },
      14: {
        fontSize: '$14',
        letterSpacing: '-.016em',
      },
      18: {
        fontSize: '$18',
      },
      24: {
        fontSize: '$24',
      },
      36: {
        fontSize: '$36',
      },
      48: {
        fontSize: '48px',
      },
      64: {
        fontSize: '64px',
      },
      72: {
        fontSize: '72px',
      },
      96: {
        fontSize: '96px',
      },
      144: {
        fontSize: '144px',
      },
      288: {
        fontSize: '288px',
      },
    },
    weight: {
      bold: {
        fontWeight: '$bold',
      },
      semiBold: {
        fontWeight: '$semiBold',
      },
      medium: {
        fontWeight: '$medium',
      },
      regular: {
        fontWeight: '$regular',
      },
    },
  },
});
