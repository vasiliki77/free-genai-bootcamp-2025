
import { NavLink } from "react-router-dom";

const NavBar = () => {
  const routes = [
    { path: "/dashboard", label: "Dashboard" },
    { path: "/study-activities", label: "Study Activities" },
    { path: "/words", label: "Words" },
    { path: "/groups", label: "Word Groups" },
    { path: "/sessions", label: "Sessions" },
    { path: "/settings", label: "Settings" },
  ];

  return (
    <nav className="fixed top-0 left-0 right-0 h-16 border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-full items-center">
        <div className="mr-4 hidden md:flex">
          <h2 className="text-xl font-bold">Learning Ancient Greek</h2>
        </div>
        <div className="flex flex-1 items-center justify-between space-x-2 md:justify-end">
          <div className="flex items-center space-x-4">
            {routes.map((route) => (
              <NavLink
                key={route.path}
                to={route.path}
                className={({ isActive }) =>
                  `nav-link ${isActive ? "active" : ""}`
                }
              >
                {route.label}
              </NavLink>
            ))}
          </div>
        </div>
      </div>
    </nav>
  );
};

export default NavBar;
