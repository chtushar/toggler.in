import { useFormik } from 'formik';

import { Flex, Button, TextField, Label } from '../ds';
import useAuth from '../requests/useAuth';
import { signInFormSchema } from '../utils/validation';

const SignInForm = () => {
  const {
    signIn: { mutate, isLoading },
  } = useAuth();
  const { handleSubmit, handleChange, values, isValid, dirty } = useFormik({
    initialValues: {
      email: '',
      password: '',
    },
    onSubmit: (values) => {
      mutate(values);
    },
    validationSchema: signInFormSchema,
  });

  const isSubmitButtonDisabled = !isValid || !dirty || isLoading;

  return (
    <form onSubmit={handleSubmit}>
      <Flex direction='column' gap='8'>
        <Flex direction='column' gap='2'>
          <Label size={18} weight='semiBold' htmlFor='email'>
            Work Email:
          </Label>
          <TextField
            type='email'
            name='email'
            placeholder='Your work email'
            onChange={handleChange}
            value={values.email}
          />
        </Flex>
        <Flex direction='column' gap='2'>
          <Label size={18} weight='semiBold' htmlFor='password'>
            Password
          </Label>
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
          Sign in to toggler.in
        </Button>
      </Flex>
    </form>
  );
};

export default SignInForm;
