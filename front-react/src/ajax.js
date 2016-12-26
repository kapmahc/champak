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


export const post = (path, body) => {
  return fetch(
    api(path),
    {
      method: 'post',
      mode: 'cors',
      headers: {
        // 'Authorization': 'Basic ' + btoa(authHeader),
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      credentials: 'include',
      body: body,
    })
    .then(parse)
}
