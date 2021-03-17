package adapter.pattern;

/**
 * 创建实现了 AdvancedMediaPlayer 接口的实体类。
 */
public class VlcPlayer implements AdvancedMediaPlayer {
    @Override
    public void playVlc(String fileName) {
        System.out.println("playVlc: " + fileName);
    }

    @Override
    public void playMp4(String fileName) {
// do nth
    }
}
