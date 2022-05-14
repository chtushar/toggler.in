interface Environment {
  NODE_ENV: string;
  BASE_URL: string;
  API_ROOT: string;
}

const ENV: Environment = {
  NODE_ENV: process.env.NODE_ENV || '',
  BASE_URL: process.env.NEXT_PUBLIC_BASE_URL || '',
  API_ROOT: process.env.NEXT_PUBLIC_API_ROOT || '',
};

export { ENV };
