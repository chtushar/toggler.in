import { styled, keyframes } from '../stitches.config';
import * as DialogPrimitive from '@radix-ui/react-dialog';

const overlayShow = keyframes({
  '0%': { opacity: 0 },
  '100%': { opacity: 1 },
});

const contentShow = keyframes({
  '0%': { opacity: 0, transform: 'translate(-50%, -48%) scale(.96)' },
  '100%': { opacity: 1, transform: 'translate(-50%, -50%) scale(1)' },
});

const Overlay = styled(DialogPrimitive.Overlay, {
  backgroundColor: '$blackA9',
  position: 'fixed',
  inset: 0,
  '@media (prefers-reduced-motion: no-preference)': {
    animation: `${overlayShow} 150ms cubic-bezier(0.16, 1, 0.3, 1)`,
  },
});

const Root = styled(DialogPrimitive.Root);
const Portal = styled(DialogPrimitive.Portal);
const Content = styled(DialogPrimitive.Content, {
  backgroundColor: '$slate1',
  padding: '$8',
  borderRadius: '$space$4',
  boxShadow:
    'hsl(206 22% 7% / 35%) 0px 10px 38px -10px, hsl(206 22% 7% / 20%) 0px 10px 20px -15px',
  position: 'fixed',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
});
const Trigger = styled(DialogPrimitive.Trigger);
const Title = styled(DialogPrimitive.Title, {
  fontSize: '$24',
  fontWeight: '$semiBold',
});
const Close = styled(DialogPrimitive.Close);

export { Overlay, Content, Trigger, Portal, Title, Root, Close };
