import { styled } from '../stitches.config';

export const Button = styled('button', {
  cursor: 'pointer',
  padding: '$3 $4',
  color: '$slate1',
  border: '1px solid transparent',
  borderRadius: '$space$4',
  display: 'flex',
  flexDirection: 'row',
  justifyContent: 'center',
  alignItems: 'center',
  gap: '$4',
  fontSize: '$18',
  outline: 'none',

  '&:focus': {
    border: '1px solid $colors$blue9',
    boxShadow: '$active-blue',
  },

  variants: {
    appearance: {
      primary: {
        fontWeight: '$semiBold',
        backgroundColor: '$slate12',
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
        '&:disabled': {
          border: 'none',
          cursor: 'default',
          color: '$slate4',
          backgroundColor: '$slate8',
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
        justifyContent: 'flex-start',
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
        padding: '$3 $4',
        backgroundColor: '$slate3',
        color: '$slate12',
        fontSize: '$14',
        justifyContent: 'space-between',
        outline: 'none',
        '&:active,&:focus': {
          backgroundColor: '$slate1',
          border: '1px solid $colors$blue9',
          boxShadow: '$active-blue',
        },
      },
      transparent: {
        padding: '$4',
        fontWeight: '$bold',
        backgroundColor: 'transparent',
        color: '$slate12',
        justifyContent: 'flex-start',
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
