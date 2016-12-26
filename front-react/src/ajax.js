export const api = (path) => {
  return `${process.env.REACT_APP_BACKEND}${path}`
}

export const get = (path) => {
  return fetch(api(path), { method: 'get' })
    .then(function(response) {
      return response.json()
    })
}
