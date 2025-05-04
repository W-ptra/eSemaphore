function Navbar() {
    return (
        <nav
            className="md:w-15 bg-white flex flex-row w-full md:flex-col items-center py-2 z-20"
            style={{ boxShadow: '2px 0 4px rgba(0, 0, 0, 0.1)' }}
        >
            <ul
                className="flex-1/2 flex flex-row md:flex-col gap-y-3 justify-evenly md:justify-start"
            >
                <li className="hover:bg-gray-100 p-1 rounded-xl cursor-pointer flex items-center justify-center" >

                    <img
                        src="/logo/chat_logo.png" alt=""
                        className="size-9"
                    />
                </li>
                <a href="/new-chat">
                    <li className="hover:bg-gray-100 p-1 rounded-xl cursor-pointer flex items-center justify-center" >

                        <img
                            src="/logo/add_chat_logo.png" alt=""
                            className="size-7 ml-1"
                        />
                    </li>
                </a>
                <li 
                    className="md:hidden hover:bg-gray-100 p-1 rounded-xl cursor-pointer flex items-center justify-center"
                >
                    <img
                        src="/logo/account_logo.png" alt=""
                        className="size-8"
                    />
                </li>
            </ul>
            <ul
                className="hidden flex-1/2 justify-end md:flex flex-col gap-y-3"
            >
                <a href="/profile">
                    <li 
                        className="hover:bg-gray-100 p-1 rounded-xl cursor-pointer flex items-center justify-center"
                    >
                        <img
                            src="/logo/account_logo.png" alt=""
                            className="size-9"
                        />
                    </li>
                </a>
            </ul>
        </nav>
    )
}

export default Navbar