import { useFormik } from 'formik';

import { Flex, Button } from '../ds';
import useCreateUser from '../mutations/useCreateUser';
import { signupFormSchema } from '../utils/validation';

const SignupForm = () => {
  const { handleCreateUser } = useCreateUser();
  const { handleChange, handleSubmit, values, isValid, dirty } = useFormik({
    initialValues: {
      firstName: '',
      lastName: '',
      email: '',
      password: '',
    },
    validationSchema: signupFormSchema,
    onSubmit: handleCreateUser,
  });

  const isSubmitButtonDisabled = !isValid || !dirty;

  return (
    <form onSubmit={handleSubmit}>
      <Flex direction='column' gap='8'>
        <label htmlFor='firstName'>
          First Name:
          <input
            type='text'
            name='firstName'
            onChange={handleChange}
            value={values.firstName}
          />
        </label>
        <label htmlFor='lastName'>
          Last Name:
          <input
            type='text'
            name='lastName'
            onChange={handleChange}
            value={values.lastName}
          />
        </label>
        <label htmlFor='email'>
          Email:
          <input
            type='email'
            name='email'
            onChange={handleChange}
            value={values.email}
          />
        </label>
        <label htmlFor='password'>
          Password
          <input
            type='password'
            name='password'
            onChange={handleChange}
            value={values.password}
          />
        </label>
        <Button type='submit' disabled={isSubmitButtonDisabled}>
          Create free account
        </Button>
      </Flex>
    </form>
  );
};

export default SignupForm;
