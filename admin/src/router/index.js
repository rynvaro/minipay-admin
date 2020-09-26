import Vue from 'vue'
import Router from 'vue-router'
import plates from '@/components/plates'
import coupons from '@/components/coupons'
import pubstores from '@/components/pubstores'
import stores from '@/components/stores'
import users from '@/components/users'
import events from '@/components/events'
import finance from '@/components/finance'
import withdraws from '@/components/withdraws'
import others from '@/components/others'
import qrcodes from '@/components/qrcodes'


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
    },
    {
      path: '/components/events',
      name: 'events',
      component: events
    },
    {
      path: '/components/finance',
      name: 'finance',
      component: finance
    },
    {
      path: '/components/withdraws',
      name: 'withdraws',
      component: withdraws
    },
    {
      path: '/components/others',
      name: 'others',
      component: others
    },
    {
      path: '/components/qrcodes',
      name: 'qrcodes',
      component: qrcodes
    }
  ]
})
