import React from 'react';
import { styled } from '../stitches.config';

const config = {
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
    color: {
      slate12: {
        color: '$slate12',
      },
      slate10: {
        color: '$slate10',
      },
    },
    weight: {
      bold: {
        fontWeight: '$bold',
      },
      medium: {
        fontWeight: '$medium',
      },
      semiBold: {
        fontWeight: '$semiBold',
      },
      regular: {
        fontWeight: '$regular',
      },
    },
  },
};

const StyledText = styled('span', config);

export interface TextProps {
  as?: React.ElementType;
}
export const Text = (props: {
  as: React.ElementType;
  children: React.ReactNode;
  size?: keyof typeof config['variants']['size'];
  color?: keyof typeof config['variants']['color'];
  weight?: keyof typeof config['variants']['weight'];
}) => {
  const { as = 'div', children, ...rest } = props;
  return (
    <StyledText as={as} {...rest}>
      {children}
    </StyledText>
  );
};
export const Label = styled('label', config);
