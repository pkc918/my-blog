import { RouteRecordRaw } from "vue-router";
const Layout = () => import("@/components/Layout.vue");

const routes: RouteRecordRaw[] = [
  {
    path: "/:catchAll(.*)",
    name: "404",
    component: Layout,
  },
  {
    path: "/",
    name: "root",
    component: Layout,
    redirect: { name: "intro" },
    children: [
      {
        path: "/intro",
        name: "intro",
        component: Layout,
      },
      {
        path: "/blog",
        name: "blog",
        component: Layout,
      },
      {
        path: "/projects",
        name: "projects",
        component: Layout,
      },
      {
        path: "/photo",
        name: "photo",
        component: Layout,
      },
      {
        path: "/music",
        name: "music",
        component: Layout,
      },
    ],
  },
];

export default routes;
