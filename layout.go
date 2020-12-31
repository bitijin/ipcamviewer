package main
var layout = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.38.2 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <object class="GtkApplicationWindow" id="addCamera">
    <property name="can-focus">False</property>
    <property name="title" translatable="yes">libvlc-go media player</property>
    <property name="window-position">center</property>
    <property name="default-width">300</property>
    <property name="default-height">300</property>
    <property name="icon-name">applications-multimedia</property>
    <property name="type-hint">dialog</property>
    <child>
      <object class="GtkBox" id="windowBox1">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="orientation">vertical</property>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <!-- n-columns=2 n-rows=1 -->
              <object class="GtkGrid">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <child>
                  <object class="GtkLabel">
                    <property name="width-request">100</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">False</property>
                    <property name="label" translatable="yes">Name:</property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left-attach">0</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkEntry" id="newName">
                    <property name="width-request">172</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">True</property>
                  </object>
                  <packing>
                    <property name="left-attach">1</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <!-- n-columns=2 n-rows=1 -->
              <object class="GtkGrid">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <child>
                  <object class="GtkLabel">
                    <property name="width-request">100</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">False</property>
                    <property name="label" translatable="yes">IP Address:</property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left-attach">0</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkEntry" id="newAddress">
                    <property name="width-request">172</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">True</property>
                  </object>
                  <packing>
                    <property name="left-attach">1</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <!-- n-columns=2 n-rows=1 -->
              <object class="GtkGrid">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <child>
                  <object class="GtkLabel">
                    <property name="width-request">100</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">False</property>
                    <property name="label" translatable="yes">Stream:</property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left-attach">0</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkEntry" id="newStream">
                    <property name="width-request">172</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">True</property>
                  </object>
                  <packing>
                    <property name="left-attach">1</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <!-- n-columns=2 n-rows=1 -->
              <object class="GtkGrid">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <child>
                  <object class="GtkSpinButton" id="cameraPosition">
                    <property name="visible">True</property>
                    <property name="can-focus">True</property>
                    <property name="text" translatable="yes">0</property>
                    <property name="adjustment">cameraPositionAdjuster</property>
                    <property name="climb-rate">1</property>
                    <property name="numeric">True</property>
                  </object>
                  <packing>
                    <property name="left-attach">1</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkLabel">
                    <property name="visible">True</property>
                    <property name="can-focus">False</property>
                    <property name="label" translatable="yes">   Position:   </property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left-attach">0</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="pack-type">end</property>
                <property name="position">3</property>
              </packing>
            </child>
            <child>
              <!-- n-columns=2 n-rows=1 -->
              <object class="GtkGrid">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <child>
                  <object class="GtkLabel">
                    <property name="width-request">100</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">False</property>
                    <property name="label" translatable="yes">Username:</property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left-attach">0</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkEntry" id="newUsername">
                    <property name="width-request">172</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">True</property>
                  </object>
                  <packing>
                    <property name="left-attach">1</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">4</property>
              </packing>
            </child>
            <child>
              <!-- n-columns=2 n-rows=1 -->
              <object class="GtkGrid">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <child>
                  <object class="GtkLabel">
                    <property name="width-request">100</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">False</property>
                    <property name="label" translatable="yes">Password:</property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left-attach">0</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkEntry" id="newPassword">
                    <property name="width-request">172</property>
                    <property name="height-request">80</property>
                    <property name="visible">True</property>
                    <property name="can-focus">True</property>
                    <property name="visibility">False</property>
                    <property name="invisible-char">●</property>
                    <property name="input-purpose">password</property>
                  </object>
                  <packing>
                    <property name="left-attach">1</property>
                    <property name="top-attach">0</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">5</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkButtonBox" id="videoControls1">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
            <property name="layout-style">end</property>
            <child>
              <object class="GtkButton" id="closeAddCameraButton">
                <property name="label">gtk-close</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <property name="use-stock">True</property>
                <property name="always-show-image">True</property>
                <signal name="clicked" handler="onClickCloseAddCamera" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="pack-type">end</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="addCameraButton">
                <property name="label">gtk-add</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <property name="use-stock">True</property>
                <property name="always-show-image">True</property>
                <signal name="clicked" handler="onClickConfirmAddCamera" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="pack-type">end</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="pack-type">end</property>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="titlebar">
      <object class="GtkHeaderBar" id="appMenuHeader2">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="title" translatable="yes">Add Camera</property>
        <child>
          <object class="GtkMenuBar" id="appMenu2">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
          </object>
          <packing>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkApplicationWindow" id="appWindow">
    <property name="visible">True</property>
    <property name="can-focus">False</property>
    <property name="has-focus">True</property>
    <property name="title" translatable="yes">libvlc-go media player</property>
    <property name="window-position">center</property>
    <property name="default-width">1280</property>
    <property name="default-height">720</property>
    <property name="icon-name">applications-multimedia</property>
    <property name="type-hint">dialog</property>
    <property name="has-resize-grip">True</property>
    <child>
      <object class="GtkBox" id="windowBox">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="orientation">vertical</property>
        <child>
          <object class="GtkDrawingArea" id="playerArea">
            <property name="sensitive">False</property>
            <property name="can-focus">True</property>
            <signal name="draw" handler="onDrawPlayerArea" swapped="no"/>
            <signal name="realize" handler="onRealizePlayerArea" swapped="no"/>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="titlebar">
      <object class="GtkHeaderBar" id="appMenuHeader">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="title" translatable="yes">IP Cam Viewer</property>
        <property name="subtitle" translatable="yes">Alpha 0.0.1</property>
        <property name="show-close-button">True</property>
        <child>
          <object class="GtkButton" id="playButton">
            <property name="label">gtk-media-play</property>
            <property name="visible">True</property>
            <property name="can-focus">True</property>
            <property name="receives-default">True</property>
            <property name="use-stock">True</property>
            <property name="always-show-image">True</property>
            <signal name="clicked" handler="onClickPlayButton" swapped="no"/>
          </object>
          <packing>
            <property name="pack-type">end</property>
          </packing>
        </child>
        <child>
          <object class="GtkEntry" id="urlInput">
            <property name="visible">True</property>
            <property name="can-focus">True</property>
          </object>
          <packing>
            <property name="pack-type">end</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkMenuBar" id="appMenu">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
          </object>
          <packing>
            <property name="position">2</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="cameraListButton">
            <property name="label" translatable="yes">Cameras</property>
            <property name="visible">True</property>
            <property name="can-focus">True</property>
            <property name="receives-default">True</property>
            <signal name="clicked" handler="onClickCameraList" swapped="no"/>
          </object>
          <packing>
            <property name="position">3</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkApplicationWindow" id="cameraListWindow">
    <property name="can-focus">False</property>
    <property name="title" translatable="yes">libvlc-go media player</property>
    <property name="window-position">center</property>
    <property name="default-width">800</property>
    <property name="default-height">400</property>
    <property name="icon-name">applications-multimedia</property>
    <property name="type-hint">dialog</property>
    <child>
      <object class="GtkBox" id="windowBox2">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="orientation">vertical</property>
        <child>
          <object class="GtkTable">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
            <property name="n-rows">2</property>
            <property name="n-columns">4</property>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="label" translatable="yes">Name</property>
                <attributes>
                  <attribute name="weight" value="bold"/>
                  <attribute name="scale" value="3"/>
                </attributes>
              </object>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="label" translatable="yes">IP Address</property>
                <attributes>
                  <attribute name="weight" value="bold"/>
                  <attribute name="scale" value="3"/>
                </attributes>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="right-attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="label" translatable="yes">Stream</property>
                <attributes>
                  <attribute name="weight" value="bold"/>
                  <attribute name="scale" value="3"/>
                </attributes>
              </object>
              <packing>
                <property name="left-attach">2</property>
                <property name="right-attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="name0">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="top-attach">1</property>
                <property name="bottom-attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="name1">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="top-attach">2</property>
                <property name="bottom-attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="name2">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="top-attach">3</property>
                <property name="bottom-attach">4</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="name4">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="top-attach">5</property>
                <property name="bottom-attach">6</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="name3">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="top-attach">4</property>
                <property name="bottom-attach">5</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="address0">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="right-attach">2</property>
                <property name="top-attach">1</property>
                <property name="bottom-attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="address1">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="right-attach">2</property>
                <property name="top-attach">2</property>
                <property name="bottom-attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="address4">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="right-attach">2</property>
                <property name="top-attach">5</property>
                <property name="bottom-attach">6</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="address3">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="right-attach">2</property>
                <property name="top-attach">4</property>
                <property name="bottom-attach">5</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="address2">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="right-attach">2</property>
                <property name="top-attach">3</property>
                <property name="bottom-attach">4</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="stream0">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">2</property>
                <property name="right-attach">3</property>
                <property name="top-attach">1</property>
                <property name="bottom-attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="stream1">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">2</property>
                <property name="right-attach">3</property>
                <property name="top-attach">2</property>
                <property name="bottom-attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="stream2">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">2</property>
                <property name="right-attach">3</property>
                <property name="top-attach">3</property>
                <property name="bottom-attach">4</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="stream3">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">2</property>
                <property name="right-attach">3</property>
                <property name="top-attach">4</property>
                <property name="bottom-attach">5</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="stream4">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
              </object>
              <packing>
                <property name="left-attach">2</property>
                <property name="right-attach">3</property>
                <property name="top-attach">5</property>
                <property name="bottom-attach">6</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="load0">
                <property name="label" translatable="yes">View</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <signal name="clicked" handler="onClickView0" swapped="no"/>
              </object>
              <packing>
                <property name="left-attach">3</property>
                <property name="right-attach">4</property>
                <property name="top-attach">1</property>
                <property name="bottom-attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="load1">
                <property name="label" translatable="yes">View</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <signal name="clicked" handler="onClickView1" swapped="no"/>
              </object>
              <packing>
                <property name="left-attach">3</property>
                <property name="right-attach">4</property>
                <property name="top-attach">2</property>
                <property name="bottom-attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="load2">
                <property name="label" translatable="yes">View</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <signal name="clicked" handler="onClickView2" swapped="no"/>
              </object>
              <packing>
                <property name="left-attach">3</property>
                <property name="right-attach">4</property>
                <property name="top-attach">3</property>
                <property name="bottom-attach">4</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="load3">
                <property name="label" translatable="yes">View</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <signal name="clicked" handler="onClickView3" swapped="no"/>
              </object>
              <packing>
                <property name="left-attach">3</property>
                <property name="right-attach">4</property>
                <property name="top-attach">4</property>
                <property name="bottom-attach">5</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="load4">
                <property name="label" translatable="yes">View</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <signal name="clicked" handler="onClickView4" swapped="no"/>
              </object>
              <packing>
                <property name="left-attach">3</property>
                <property name="right-attach">4</property>
                <property name="top-attach">5</property>
                <property name="bottom-attach">6</property>
              </packing>
            </child>
            <child>
              <placeholder/>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkButtonBox" id="videoControls2">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
            <property name="layout-style">end</property>
            <child>
              <object class="GtkButton" id="closeCameraListButton">
                <property name="label">gtk-close</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <property name="use-stock">True</property>
                <property name="always-show-image">True</property>
                <signal name="clicked" handler="onClickCloseCameraList" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="pack-type">end</property>
                <property name="position">0</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="pack-type">end</property>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="titlebar">
      <object class="GtkHeaderBar" id="appMenuHeader1">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="title" translatable="yes">Camera List</property>
        <child>
          <object class="GtkMenuBar" id="appMenu1">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
          </object>
          <packing>
            <property name="position">2</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="openAddCameraButton">
            <property name="label">gtk-add</property>
            <property name="visible">True</property>
            <property name="can-focus">True</property>
            <property name="receives-default">True</property>
            <property name="use-stock">True</property>
            <property name="always-show-image">True</property>
            <signal name="clicked" handler="onClickAddCamera" swapped="no"/>
          </object>
          <packing>
            <property name="pack-type">end</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkAdjustment" id="cameraPositionAdjuster">
    <property name="upper">5</property>
    <property name="value">1</property>
    <property name="step-increment">1</property>
    <property name="page-increment">4</property>
  </object>
</interface>
`