import { styled } from '../stitches.config';

export const TextField = styled('input', {
  // Reset
  appearance: 'none',
  boxSizing: 'border-box',
  fontFamily: 'inherit',
  margin: '0',
  outline: 'none',
  width: '100%',
  WebkitTapHighlightColor: 'rgba(0,0,0,0)',
  '&::before': {
    boxSizing: 'border-box',
  },
  '&::after': {
    boxSizing: 'border-box',
  },

  // Custom
  border: '1px solid $gray11',
  padding: '$3',
  borderRadius: '$space$1',
  '&::placeholder': {
    color: '$gray8',
  },
});
