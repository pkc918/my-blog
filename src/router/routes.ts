import { RouteRecordRaw } from "vue-router";
const Layout = () => import("@/components/Layout.vue");
const Intro = () => import("@/views/Intro.vue");
const Blogs = () => import("@/views/Blogs.vue");
const Projects = () => import("@/views/Projects.vue");
const Photo = () => import("@/views/Photo.vue");
const Music = () => import("@/views/Music.vue");

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
        component: Intro,
      },
      {
        path: "/blog",
        name: "blog",
        component: Blogs,
      },
      {
        path: "/projects",
        name: "projects",
        component: Projects,
      },
      {
        path: "/photo",
        name: "photo",
        component: Photo,
      },
      {
        path: "/music",
        name: "music",
        component: Music,
      },
    ],
  },
];

export default routes;
