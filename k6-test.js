import http from 'k6/http';

import { sleep } from 'k6';

export const options = {
  vus: 100,
  iterations: 100,
};

export default function () {
  const payload = JSON.stringify({
    character: {
      name: 'Bruce',
    },
    movie: {
      title: 'Batman Begins',
      year: 2005,
    },
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  // POST request
  http.post('http://localhost:8080/movie', payload, params);

  // Put request
  http.put('http://localhost:8080/movie/0', payload, params);

  // Get request
  http.get('http://localhost:8080/movie');

  // Delete request
  http.del('http://localhost:8080/movie/0');
  sleep(1);
}
