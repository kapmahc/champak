const items = (user) => {
  return [
    {
      label: 'auth.profile',
      links: [
        {label: 'auth.users.logs.title', href:'/users/logs'},
        {label: 'auth.users.profile.title', href:'/users/profile'},
        {label: 'auth.users.change-password.title', href:'/users/change-password'}
      ]
    }
  ]
}

export default items
