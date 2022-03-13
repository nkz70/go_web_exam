import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ScoreView from "../views/ScoreView.vue";
import UserCreateView from "../views/UserCreate.vue"

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
  ],
});

export default router;
