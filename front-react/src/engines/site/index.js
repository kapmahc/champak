import routes from './routes'

export default {
  dashboard(user){
    var links = []
    if(user.is_admin){
      links.push({
        label: 'site.profile',
        links: [
          {label: 'site.info.title', href:'/site/info'},
          {label: 'site.author.title', href:'/site/author'},
          {label: 'site.seo.title', href:'/site/seo'},
          {label: 'site.smtp.title', href:'/site/smtp'},
          {label: 'site.status.title', href:'/site/status'},
          {label: 'site.leave-words.title', href:'/site/leave-words'},
          {label: 'site.notices.title', href:'/site/notices'},
          {label: 'site.links.title', href:'/site/links'},
          {label: 'site.cards.title', href:'/site/cards'},
          {label: 'site.locales.title', href:'/site/locales'},
          {label: 'site.users.title', href:'/site/users'}
        ]
      })
    }
    return links
  },
  routes
}
