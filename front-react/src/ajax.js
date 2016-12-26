export const api = (path) => {
  return `${process.env.REACT_APP_BACKEND}${path}`
}

export const get = (path, success) => {
  return fetch(api(path), { method: 'get', credentials: 'include'  })
    .then(function(response) {
      return response.json()
    }).then(success)
}
