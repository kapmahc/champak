const items = (user) => {
  var links = []
  if(user.admin){
    links.push({
      label: 'site.profile',
      links: [
        {label: 'site.info.title', href:'/admin/site/info'},
        {label: 'site.author.title', href:'/admin/site/author'},
        {label: 'site.seo.title', href:'/admin/site/seo'},
        {label: 'site.smtp.title', href:'/admin/site/smtp'},
        {label: 'site.status.title', href:'/admin/site/status'},
        {label: 'site.leave-words.index.title', href:'/leave-words'},
        {label: 'site.notices.title', href:'/notices'},
        {label: 'site.links.title', href:'/links'},
        {label: 'site.cards.title', href:'/cards'},
        {label: 'site.locales.title', href:'/locales'},
        {label: 'site.users.title', href:'/users'}
      ]
    })
  }
  return links
}

export default items
