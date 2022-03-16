import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ScoreView from "../views/ScoreView.vue";
import UserListView from "../views/UserListView.vue"
import UserCreateView from "../views/UserCreateView.vue"
import UserEditView from "../views/UserEditView.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Home",
      component: HomeView,
    },
    {
      path: "/user/:id/score",
      name: "Score",
      component: ScoreView,
    },
    {
      path: "/users",
      name: "UserList",
      component: UserListView,
    },
    {
      path: "/user",
      name: "UserCreate",
      component: UserCreateView,
    },
    {
      path: "/user/:id",
      name: "UserEdit",
      component: UserEditView,
    },
  ],
});

export default router;
