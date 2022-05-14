import { useFormik } from 'formik';

import { Flex, Button, TextField, Text } from '../ds';
import useCreateUser from '../requests/useCreateUser';
import { signupFormSchema } from '../utils/validation';

const SignupForm = () => {
  const { mutate, isLoading } = useCreateUser();
  const { handleChange, handleSubmit, values, isValid, dirty } = useFormik({
    initialValues: {
      name: '',
      email: '',
      password: '',
    },
    validationSchema: signupFormSchema,
    onSubmit: (values) => {
      mutate(values);
    },
  });

  const isSubmitButtonDisabled = !isValid || !dirty || isLoading;

  return (
    <form onSubmit={handleSubmit}>
      <Flex direction='column' gap='8'>
        <Flex direction='column' gap='2'>
          <Text as='label' size={18} weight='semiBold' htmlFor='name'>
            Name:
          </Text>
          <TextField
            type='text'
            name='name'
            placeholder='Your name'
            onChange={handleChange}
            value={values.name}
          />
        </Flex>
        <Flex direction='column' gap='2'>
          <Text as='label' size={18} weight='semiBold' htmlFor='email'>
            Work Email:
          </Text>
          <TextField
            type='email'
            name='email'
            placeholder='Your work email'
            onChange={handleChange}
            value={values.email}
          />
        </Flex>
        <Flex direction='column' gap='2'>
          <Text as='label' size={18} weight='semiBold' htmlFor='password'>
            Password
          </Text>
          <TextField
            type='password'
            name='password'
            placeholder='*****'
            onChange={handleChange}
            value={values.password}
          />
        </Flex>
        <Button
          appearance='primary'
          type='submit'
          disabled={isSubmitButtonDisabled}
        >
          Create free account
        </Button>
      </Flex>
    </form>
  );
};

export default SignupForm;
