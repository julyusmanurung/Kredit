import React, {useEffect, useState} from "react";
import {Image} from "react-bootstrap"
import './sidebar.component.css'
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {
    faBars,
    faExchange,
    faFile,
    faList,
    faMoon,
    faSignOut,
    faSun,
    faDashboard
} from '@fortawesome/free-solid-svg-icons'
import LogoSinarmas from '../../assets/logo-only.png'
import Swal from "sweetalert2";
import {Link, useLocation} from "react-router-dom";

export const SideBar = () => {
    const location = useLocation();
    const [activePage, setActivePage] = useState(location.pathname.substring(1));
    const toggleSidebar = () => {
        const nav = document.getElementById("nav")
        nav.classList.toggle("close")
        if (localStorage.getItem("sidebar") === "open") {
            localStorage.setItem("sidebar", "close")
        } else {
            localStorage.setItem("sidebar", "open")
        }
    }
    const handleSignOut = () => {
        Swal.fire({
            title: 'You are about to sign out',
            text: "Are you sure?",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes'
        }).then((result) => {
            if (result.isConfirmed) {
                localStorage.removeItem("user")
                localStorage.removeItem("mode")
                localStorage.removeItem("nik")
                localStorage.removeItem("sidebar")
                Swal.fire({
                    icon: 'success',
                    title: 'You already signed out',
                    showConfirmButton: false,
                    timer: 1500
                }).then(() => window.location.href = "/")
            }
        })
    }

    const toggleMode = () => {
        if (localStorage.getItem("mode") === "dark") {
            localStorage.setItem("mode", "light")
        } else {
            localStorage.setItem("mode", "dark")
        }

        const body = document.getElementById("body")
        body.classList.toggle("dark")
    }

    const toggleSidebarSubMenu = () => {
        const sideBarMenuTrigger = document.getElementById("sub-menu-1-trigger")
        const sideBarMenu = document.getElementById("sub-menu-1")

        sideBarMenuTrigger.classList.toggle('closed')
        sideBarMenu.classList.toggle('closed')
    }

    return (
        <nav className={localStorage.getItem("sidebar") === "open" ? "sidebar" : "sidebar close"} id="nav">
            <header>
                <div className="image-text">
                <span className="image">
                    <Image src={LogoSinarmas}></Image>
                </span>
                <div className="text logo-text">
                    <span className="name">SI Pengajuan Kredit</span>
                    <span className="profession">Bank Sinarmas</span>
                </div>
                </div>
                <i className='toggle' id="toggle" onClick={() => {toggleSidebar()}}>
                        <FontAwesomeIcon icon={faBars}/>
                </i>
            </header>
            <div className="menu-bar">
                <div className="menu">
                    <ul className="menu-links">
                        <li className={activePage === "" ? "nav-link active-page": "nav-link"} onClick={() => {
                            setActivePage("")
                            document.getElementById("sub-menu-1").classList.add("closed");
                            document.getElementById("sub-menu-1-trigger").classList.add("closed");
                        }}>
                            <Link to="/">
                                <i className={activePage === "" ? "icon active-page" : "icon"}><FontAwesomeIcon icon={faDashboard}/></i>
                                <span className={activePage === "" ? "text nav-text active-page" : "text nav-text"}>Dasbor</span>
                            </Link>
                        </li>
                        <li onClick={toggleSidebarSubMenu} id="sub-menu-1-trigger" className={activePage === "checklist" ? "dropdown active-page closed" : "closed dropdown"}>
                            <Link to="#" className="nav-link dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                <i className={activePage === "checklist" ? "active-page icon" : "icon"}><FontAwesomeIcon icon={faExchange}/></i>
                                <span className={activePage === "checklist" ? "text nav-text active-page" : "text nav-text"}>Transaksi</span>
                            </Link>
                        </li>
                        <div id="sub-menu-1" className="closed">
                            <li className="nav-link" onClick={() => {
                                setActivePage("checklist")
                                document.getElementById("sub-menu-1-trigger").classList.add("closed");
                                document.getElementById("sub-menu-1").classList.add("closed");
                            }}>
                                <Link to="/checklist">
                                    <i className="icon"><FontAwesomeIcon icon={faList}/></i>
                                    <span className="text nav-text">Checklist Pencairan</span>
                                </Link>
                            </li>
                        </div>
                        <li className={activePage === "laporan" ? "nav-link active-page": "nav-link"} onClick={() => {
                            setActivePage("laporan")
                            document.getElementById("sub-menu-1-trigger").classList.add("closed");
                            document.getElementById("sub-menu-1").classList.add("closed");
                        }}>
                            <Link to="/laporan">
                                <i className={activePage === "laporan" ? "icon active-page" : "icon"}><FontAwesomeIcon icon={faFile}/></i>
                                <span className={activePage === "laporan" ? "text nav-text active-page" : "text nav-text"}>Laporan</span>
                            </Link>
                        </li>
                        <li className="mode">
                            <div className="sun-moon">
                                <i className='icon moon'><FontAwesomeIcon icon={faMoon}/></i>
                                <i className='icon sun'><FontAwesomeIcon icon={faSun}/></i>
                            </div>
                            <span className="mode-text text" id="mode-text">{localStorage.getItem("mode") === "light" ? "Mode Terang" : "Mode Gelap"}</span>
                            <div className="toggle-switch" onClick={() => {
                                toggleMode()
                                // window.location.href = window.location.pathname
                            }}>
                                <span className="switch"></span>
                            </div>
                        </li>
                    </ul>
                </div>
                <div className="bottom-content">
                    <li onClick={() => handleSignOut()}>
                        <Link to="#">
                            <i className='icon'><FontAwesomeIcon icon={faSignOut}/></i>
                            <span className="text nav-text">Logout</span>
                        </Link>
                    </li>
                </div>
            </div>
        </nav>
    )
}