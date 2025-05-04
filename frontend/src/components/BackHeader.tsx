function BackHeader() {
    return (
        <header
            className="h-15 flex items-center bg-white z-30  px-3"
            style={{ boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)' }}
        >
            <div
                className="flex-1 flex items-center h-full bg-white"

            >
                <a href="/">
                    <div
                        className="block hover:bg-gray-100 rounded"
                    >
                        <img
                            src="/logo/back_logo.png" alt=""
                            className="w-10 h-10"
                        />
                    </div>
                </a>
                <h4 className="text-blue-400 font-bold ml-1 text-xl">
                    Profile
                </h4>
            </div>
        </header>
    )
}

export default BackHeader