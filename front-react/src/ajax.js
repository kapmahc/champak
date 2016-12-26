export const api = (path) => {
  return `${process.env.REACT_APP_BACKEND}${path}`
}

export const get = (path) => {
  return fetch(api(path), { method: 'get', credentials: 'include'  })
    .then(function(response) {
      return response.status === 200 ? response.json() : response.text()
    })
}


export const post = (path, form) => {
  return fetch(api(path), { method: 'post', credentials: 'include', body: form  })
    .then(function(response) {
      return response.status === 200 ? response.json() : response.text()
    })
}
