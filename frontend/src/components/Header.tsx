function Header() {
    return (
        <header
            className="h-15 flex items-center bg-white z-30  px-3"
            style={{ boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)' }}
        >
            <div
                className="flex-1 flex items-center h-full bg-white"
                
            >
                <div
                    className="block md:hidden"
                >
                    <img
                        src="/logo/back_logo.png" alt=""
                        className="w-10 h-10"
                    />
                </div>
                <img
                    src="/logo/eSemaphore_logo_white.webp" alt=""
                    className="w-10 h-10"
                />
                <h4 className="text-blue-400 font-bold ml-1 text-xl">
                    eSemaphore
                </h4>
            </div>
            <div
                className="flex items-center justify-center cursor-pointer hover:bg-gray-100 pr-2 rounded"
            >
                <img src="/logo/logout_logo.png" alt="" 
                    className="size-10"
                />
                <p
                    className="text-blue-500"
                >
                    Log out
                </p>
            </div>
        </header>
    )
}

export default Header