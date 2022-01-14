import { useFormik } from 'formik';
import type { NextPage } from 'next';
import useCreateUser from '../mutations/useCreateUser';

const Signup: NextPage = () => {
  const { handleCreateUser } = useCreateUser();
  const { handleChange, handleSubmit, values } = useFormik({
    initialValues: {
      firstName: undefined,
      lastName: undefined,
      email: undefined,
      password: undefined,
    },
    onSubmit: handleCreateUser,
  });

  return (
    <div>
      <form onSubmit={handleSubmit}>
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
        <button type='submit'>Register</button>
      </form>
    </div>
  );
};

export default Signup;
