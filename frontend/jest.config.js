module.exports = {
    preset: 'ts-jest',
    testEnvironment: 'jsdom',
    transform: {
        '^.+\\.(ts|tsx)$': 'ts-jest', // TypeScriptファイルを変換するための設定
    },
    transformIgnorePatterns: [
        "/node_modules" // node_modules内のファイルは変換しない
    ],
    moduleNameMapper: {
        "^@/(.*)$": "<rootDir>/$1", // エイリアスをtsのものからjsのものに変更
        "\\.(css|less|scss|sass)$": "identity-obj-proxy", // cssをモックするための設定
    },
    setupFilesAfterEnv: ["<rootDir>/jest.setup.ts"], // Jestのセットアップファイル

}