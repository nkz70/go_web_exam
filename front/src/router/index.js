import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ScoreView from "../views/ScoreView.vue";
import UserCreateView from "../views/UserCreate.vue"
import UserEditView from "../views/UserEdit.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/user/:id/score",
      name: "score",
      component: ScoreView,
    },
    {
      path: "/user",
      name: "usercreate",
      component: UserCreateView,
    },
    {
      path: "/user/:id",
      name: "userupdate",
      component: UserEditView,
    },
  ],
});

export default router;
