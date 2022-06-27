import { styled } from '../stitches.config';

export const TextArea = styled('textarea', {
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
  fontSize: '$14',
  lineHeight: '27px',
  fontWeight: '$regular',
  border: '1px solid transparent',
  padding: '$3 $4',
  borderRadius: '$space$4',
  backgroundColor: '$slate3',
  color: '$slate12',
  '&::placeholder': {
    color: '$slate8',
    lineHeight: '27px',
  },

  '&:active,&:focus': {
    backgroundColor: '$slate1',
    border: '1px solid $colors$blue9',
    boxShadow: '$active-blue',
  },
});
