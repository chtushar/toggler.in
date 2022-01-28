import { Button } from './Button';
import { Icon } from './Icon';
import { useSelect } from 'downshift';
import { styled } from '../stitches.config';
import { Flex } from './Flex';

interface SelectProps {
  items: Array<Record<string, any>>;
  onChange?: any;
}

const SelectWrapper = styled('div', {
  position: 'relative',
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'stretch',
});

const List = styled(Flex, {
  position: 'absolute',
  top: '100%',
  listStyle: 'none',
  marginTop: '$2',
  borderRadius: '$space$2',
  overflow: 'hidden',
  boxShadow: '$shadow$1',
  zIndex: 10,
  width: '100%',
});

const ListItem = styled('li', {
  background: '$slate1',
  padding: '$4',
  fontSize: '$14',
  variants: {
    selected: {
      true: {
        background: '$slate3',
      },
      false: {
        background: '$slate1',
      },
    },
  },
});

export const Select = ({ items, onChange }: SelectProps) => {
  const {
    isOpen,
    highlightedIndex,
    getToggleButtonProps,
    getMenuProps,
    getItemProps,
    selectedItem,
  } = useSelect({
    items,
    itemToString: (item) => item?.value,
    onSelectedItemChange: onChange,
  });
  return (
    <SelectWrapper>
      <Button type='button' variant='select' {...getToggleButtonProps()}>
        {selectedItem?.label || 'Choose an option'}{' '}
        <Icon
          className={isOpen ? 'ri-arrow-up-s-line' : 'ri-arrow-down-s-line'}
        />
      </Button>
      <List as='ul' direction='column' align='stretch' {...getMenuProps()}>
        {isOpen &&
          items.map((item, index) => (
            <ListItem
              key={`${item?.value}-${index}`}
              selected={highlightedIndex === index}
              {...getItemProps({ item, index })}
            >
              {item?.label}
            </ListItem>
          ))}
      </List>
    </SelectWrapper>
  );
};

export default Select;
