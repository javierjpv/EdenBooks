import { Link, useNavigate } from "react-router";
import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import MenuIcon from "@mui/icons-material/Menu";
import Container from "@mui/material/Container";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Tooltip from "@mui/material/Tooltip";
import MenuItem from "@mui/material/MenuItem";
import { Book, Logout } from "@mui/icons-material";
import { useAuthStore } from "../modules/users/hooks/useAuthStore";
import { useEffect, useState } from "react";
import { IUserRequest } from "../modules/users/interfaces/IUserRequest";
import { userService } from "../modules/users/services/userService";
import { IUserResponse } from "../modules/users/interfaces/IUserResponse";

const settings = ["Profile", "Account", "Dashboard", "Logout"];

export function Navbar() {
  const { user, startLogout } = useAuthStore();
  const navigate = useNavigate();
  const [anchorElNav, setAnchorElNav] = React.useState<null | HTMLElement>(
    null
  );
  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(
    null
  );

  const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  const handleLogOut = () => {
    startLogout();
    navigate("/auth");
  };

  const [profileUserRequest, setprofileUserRequest] =
    useState<IUserRequest | null>(null);
  const [loadingProfileUser, setloadingProfileUser] = useState<boolean>(true);

  const fetchUser = async () => {
    console.log("fetching user");
    if (!user.ID) {
      setloadingProfileUser(false);
      return;
    }
    const response = await userService.GetUserById(user.ID);
    if (response.success && response.data) {
      console.log("VALOR EN FETCHING USER", response.data);
      const data: IUserResponse = response.data;

      setprofileUserRequest((prev) => ({
        ...prev,
        Name: data.Name,
        Tel: data.Tel,
        ImageURL: data.ImageURL,
      }));

      setloadingProfileUser(false);
      console.log(data);
    } else {
      console.log("Error al llamar a fetchUser");
    }
  };

  useEffect(() => {
    fetchUser();
  }, []);
  
  let pages =
    user?.userState === "AUTHENTICATED"
      ? [
          { name: "Mis libros", link: "/products" },
          { name: "Vender", link: "/products/new" },
          { name: "Notificaciones", link: "/notifications" },
          { name: "Pedidos", link: "/orders" },
          { name: "Chats", link: "/chats" },
          { name: "Wishlist", link: "/favorites" },
          { name: "Profile", link: "/myProfile" },
        ]
      : [{ name: "Auth", link: "/auth" }];

  return (
    <AppBar position="fixed">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <Book sx={{ display: { xs: "none", md: "flex" }, mr: 1 }} />
          <Typography
            variant="h6"
            noWrap
            component={Link}
            to="/products"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontFamily: "monospace",
              fontWeight: 700,
              letterSpacing: ".3rem",
              color: "inherit",
              textDecoration: "none",
            }}
          >
            EdenBooks
          </Typography>

          <Box sx={{ flexGrow: 1, display: { xs: "flex", md: "none" } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: "bottom",
                horizontal: "left",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "left",
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{ display: { xs: "block", md: "none" } }}
            >
              {pages.map((page) => (
                <MenuItem key={page.name} onClick={handleCloseNavMenu}>
                  <Typography
                    component={Link}
                    to={page.link}
                    sx={{ textDecoration: "none", color: "inherit" }}
                  >
                    {page.name}
                  </Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>
          <Book sx={{ display: { xs: "flex", md: "none" }, mr: 1 }} />
          <Typography
            variant="h5"
            noWrap
            component={Link}
            to="/products"
            sx={{
              mr: 2,
              display: { xs: "flex", md: "none" },
              flexGrow: 1,
              fontFamily: "monospace",
              fontWeight: 700,
              letterSpacing: ".3rem",
              color: "inherit",
              textDecoration: "none",
            }}
          >
            EdenBooks
          </Typography>
          <Box sx={{ flexGrow: 1, display: { xs: "none", md: "flex" } }}>
            {pages.map((page) => (
              <Button
                component={Link}
                to={page.link}
                key={page.name}
                onClick={handleCloseNavMenu}
                sx={{ my: 2, color: "white", display: "block" }}
              >
                {page.name}
              </Button>
            ))}
          </Box>
          {user?.userState === "AUTHENTICATED" && (
            <IconButton onClick={handleLogOut}>
              <Logout />
            </IconButton>
          )}

          <Box sx={{ flexGrow: 0 }}>
            <Tooltip title="Open settings">
              <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                {loadingProfileUser ? (
                  <>
                    <h1>Cargando</h1>
                  </>
                ) : profileUserRequest?.ImageURL ? (
                  <Avatar
                    alt="Remy Sharp"
                    src={
                      profileUserRequest.ImageURL
                    }
                  />
                ) : (
                  <Avatar
                  alt="Remy Sharp"
                  src="/static/images/avatar/2.jpg"
                  
                />
                )}
              </IconButton>
            </Tooltip>
            <Menu
              sx={{ mt: "45px" }}
              id="menu-appbar"
              anchorEl={anchorElUser}
              anchorOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              open={Boolean(anchorElUser)}
              onClose={handleCloseUserMenu}
            >
              {settings.map((setting) => (
                <MenuItem key={setting} onClick={handleCloseUserMenu}>
                  <Typography sx={{ textAlign: "center" }}>
                    {setting}
                  </Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
}
