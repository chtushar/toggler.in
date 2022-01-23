import { styled } from '../stitches.config';

export const TextField = styled('input', {
  // Reset
  appearance: 'none',
  boxSizing: 'border-box',
  fontFamily: 'inherit',
  margin: '0',
  outline: 'none',
  width: '100%',
  '&::before': {
    boxSizing: 'border-box',
  },
  '&::after': {
    boxSizing: 'border-box',
  },

  // Custom
  fontSize: '$18',
  fontWeight: '$semiBold',
  border: '1px solid transparent',
  padding: '$4',
  borderRadius: '$space$4',
  backgroundColor: '$slate3',
  color: '$slate12',
  '&::placeholder': {
    color: '$slate8',
  },

  '&:active,&:focus': {
    backgroundColor: '$slate1',
    border: '1px solid $colors$blue9',
    boxShadow: '$active-blue',
  },
});
