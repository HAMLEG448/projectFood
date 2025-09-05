import Home from '../pages/Home.tsx'
import Request from "../pages/request.tsx"
import Edit from "../pages/information/Edit.tsx"
import Add from "../pages/information/Add.tsx"

import {
  createBrowserRouter,
} from "react-router-dom";
import Sidebar from '../components/Sidebar.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Sidebar />,
    children:[
      {
        path: "/",
        element: <Home/>,
      },
      {
        path: "/home",
        element: <Home/>,
      },
      {
        path: "/request",
        element: <Request />,
      },
      {
        path: "/information",
        children:[
          {
            path: "/information/Add",
            element: <Add/>,
          },
          {
          path: "/information/Edit",
          element: <Edit/>,
          },
        ],
      },
    ],
  },
]);

export default router