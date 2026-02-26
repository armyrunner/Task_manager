import React from "react";
import {
    CDropdown,
    CDropdownToggle,
    CDropdownMenu,
    CDropdownItem,
    CAvatar,
} from '@coreui/react';

import { useNavigate } from 'react-router-dom'
import CIcon  from "@coreui/icons-react";
import { cilAccountLogout } from '@coreui/icons';
import { useAuth } from "./useAuth";


const Profile = () => {
    const navigate = useNavigate();
    const {isLoggedIn, logout, user } = useAuth();
    const userProfile = {
        name: `${user?.username}`,
        avatar: `${user?.username?.charAt(0).toLocaleUpperCase()}`,
        // or use initials fallback if no avatar
      }; 
      const handleAuthClick = () => {
        if (isLoggedIn) {
          logout();
          navigate("/home");
        } else {
          navigate("/signin");
        }
      }
    return (
        <>
            <CDropdown>
                <CDropdownToggle
                caret={false}
                className="py-0 ps-2 pe-0"
                style={{border:'none', background: 'transparent'}}
                >
                 <CAvatar color="success" size="md">
                    {userProfile.avatar}
                 </CAvatar>
                </CDropdownToggle>
                <CDropdownMenu>
                    <CDropdownItem disabled>
                        <CIcon icon={cilAccountLogout} className="me-2"/>
                        {user?.username}
                    </CDropdownItem>
                    <CDropdownItem onClick={handleAuthClick}>
                        <CIcon icon={cilAccountLogout}/>
                        Logout
                    </CDropdownItem>
                </CDropdownMenu>
            </CDropdown>
        </>
    );
};

export default Profile;