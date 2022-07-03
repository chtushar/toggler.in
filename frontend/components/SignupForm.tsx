import { useFormik } from 'formik';

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
      <div>
        <div>
          <label htmlFor='name'>Name:</label>
          <input
            type='text'
            name='name'
            placeholder='Your name'
            onChange={handleChange}
            value={values.name}
          />
        </div>
        <div>
          <label htmlFor='email'>Work Email:</label>
          <input
            type='email'
            name='email'
            placeholder='Your work email'
            onChange={handleChange}
            value={values.email}
          />
        </div>
        <label>
          <label htmlFor='password'>Password</label>
          <input
            type='password'
            name='password'
            placeholder='*****'
            onChange={handleChange}
            value={values.password}
          />
        </label>
        <button type='submit' disabled={isSubmitButtonDisabled}>
          Create free account
        </button>
      </div>
    </form>
  );
};

export default SignupForm;
