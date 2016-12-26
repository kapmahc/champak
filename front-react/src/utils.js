export const api = (p) => {
  return `${process.env.REACT_APP_BACKEND}${p}`
}

export const get = (p) => {
  return fetch(api(p), { method: 'get'})
    .then(function(response) {
      return response.json()
    })
}
