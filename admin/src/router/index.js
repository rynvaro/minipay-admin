import Vue from 'vue'
import Router from 'vue-router'
import plates from '@/components/plates'
import coupons from '@/components/coupons'
import pubstores from '@/components/pubstores'
import stores from '@/components/stores'
import users from '@/components/users'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/components/plates',
      name: 'plates',
      component: plates
    }, {
      path: '/components/coupons',
      name: 'coupons',
      component: coupons
    },
    {
      path: '/components/pubstores',
      name: 'pubstores',
      component: pubstores
    },
    {
      path: '/components/stores',
      name: 'stores',
      component: stores
    },
    {
      path: '/components/users',
      name: 'users',
      component: users
    }
  ]
})
