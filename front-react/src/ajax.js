export const api = (path) => {
  return `${process.env.REACT_APP_BACKEND}${path}`
}

export const parse = (res) => {
  return res.status === 200 || res.status === 0 ?
    res.json() :
    res.text().then(err => {throw err;})    
}

export const get = (path) => {
  return fetch(api(path), { method: 'get', credentials: 'include'  }).then(parse)
}


export const post = (path, form) => {
  return fetch(api(path), { method: 'post', credentials: 'include', body: form  }).then(parse)
}
