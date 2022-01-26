import { styled } from '../stitches.config';

export const Button = styled('button', {
  cursor: 'pointer',
  padding: '$4 $8',
  color: '$slate1',
  border: '1px solid transparent',
  borderRadius: '$space$4',
  display: 'flex',
  flexDirection: 'row',
  alignItems: 'center',
  gap: '$4',
  fontSize: '$18',

  variants: {
    appearance: {
      primary: {
        fontWeight: '$semiBold',
        backgroundColor: '$slate12',
        '&:disabled': {
          cursor: 'default',
          backgroundColor: '$slate8',
        },
        '&:hover': {
          border: '1px solid $slate12',
          color: '$slate12',
          backgroundColor: '$slate1',
        },
        '&:active': {
          border: '1px solid $slate12',
          color: '$slate12',
          backgroundColor: '$slate3',
        },
      },
      secondary: {
        fontWeight: '$semiBold',
        color: '$slate12',
        backgroundColor: 'transparent',
        '&:disabled': {
          cursor: 'default',
          backgroundColor: '$slate8',
        },
        '&:hover': {
          backgroundColor: '$blackA2',
        },
        '&:active': {
          backgroundColor: '$blackA3',
        },
      },
    },
    variant: {
      'menu-button': {
        padding: '$4',
        fontWeight: '$bold',
        backgroundColor: '$primary3',
        color: '$slate12',
        '&:hover': {
          border: '1px solid transparent',
          backgroundColor: '$primary4',
        },
        '&:active': {
          backgroundColor: '$primary6',
          border: '1px solid transparent',
        },
      },
      select: {
        padding: '$4',
        backgroundColor: '$slate3',
        color: '$slate12',
        fontSize: '$14',
      },
      transparent: {
        padding: '$4',
        fontWeight: '$bold',
        backgroundColor: 'transparent',
        color: '$slate12',
        '&:hover': {
          border: '1px solid transparent',
          backgroundColor: '$blackA2',
        },
        '&:active': {
          border: '1px solid transparent',
          backgroundColor: '$blackA3',
        },
      },
    },
  },
});
